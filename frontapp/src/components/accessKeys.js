import React, { useState } from 'react';
import {axiosInstanceWithJWT} from "../api/axios";

function AccessKeys() {

    const [role, setRole] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const [keys, setKeys] = useState(null);

    const handleRoleChange = (event) => {
        setRole(event.target.value);
    };

    const handleSubmit = async () => {

        setIsLoading(true);


        const dataToSend = {
            role:role,
        };

        try {
            const [licenseResponse] = await Promise.all([
                axiosInstanceWithJWT.post('/api/access-keys/',dataToSend),
            ]);
            setKeys(licenseResponse.data.key)
        } catch (error) {
            console.error("Ошибка при отправке данных: ", error);
        } finally {
            setIsLoading(false);
        }

    };

    return (
        <div>
            <label htmlFor="roleSelect">Select role:</label>
            <select id="roleSelect" value={role} onChange={handleRoleChange}>
                <option value="">select role</option>
                <option value="admin">admin</option>
                <option value="distributors">distributors</option>
                <option value="reseller">reseller</option>
                <option value="salesman">salesman</option>
            </select>
            <br></br><br></br>
            <button onClick={handleSubmit}>gen access_key</button>
            <h2>Keys list:</h2>
            {isLoading ? (
                <p>Loading...</p>
            ) : keys ? (
                <div>
                    <p>Key: {keys}</p>
                </div>
            ) : (
                <p></p>
            )}
        </div>
    );


}

export default AccessKeys;