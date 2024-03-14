import React, {useEffect, useState} from "react";
import {axiosInstanceWithJWT} from "../api/axios";
import {toast} from "react-toastify";

function Cheats() {
    const [cheats, setCheats] = useState([]);
    const [name, setName] = useState('');
    const [secure, setSecure] = useState('secure');
    const [isAllowedGenerate, setIsAllowedGenerate] = useState(true);

    useEffect(() => {
        const fetchUsers = async () => {
            try {
                const response = await axiosInstanceWithJWT.get('/api/cheats/');
                setCheats(response.data.cheats);
            } catch (error) {
                toast.error(`error: ${error.message}`);
            } finally {
            }
        };

        fetchUsers();
    }, []);


    const handleSecureChange = async (keyId, newValue) => {

        const updatedCheats = cheats.map(cheat =>
            cheat.id === keyId ? { ...cheat, secure: newValue } : cheat
        );

        const updatedCheat = updatedCheats.find(cheat => cheat.id === keyId);

        try {
            await axiosInstanceWithJWT.put('/api/cheats/', updatedCheat);
            setCheats(updatedCheats);
        }
        catch(error) {
            toast.error(`error: ${error.message}`);
        }
        finally {

        }
    };

    const handleAllowedGenerateChange = async (keyId, newValue) => {

        const updatedCheats = cheats.map(cheat =>
            cheat.id === keyId ? { ...cheat, is_allowed_generate: parseInt(newValue) } : cheat
        );

        const updatedCheat = updatedCheats.find(cheat => cheat.id === keyId);

        try {
            await axiosInstanceWithJWT.put('/api/cheats/', updatedCheat);
            setCheats(updatedCheats);
        }
        catch(error) {
            toast.error(`error: ${error.message}`);
        }
        finally {

        }
    };

    const handleSubmit = async (event) => {
        event.preventDefault();
        const formData = { name, secure, is_allowed_generate: isAllowedGenerate };
        try {
            var response = await axiosInstanceWithJWT.post('/api/cheats/', formData);
            const addCheat = { id :response.data.id, name : name, secure :secure, is_allowed_generate: isAllowedGenerate };
            setCheats([...cheats, addCheat]);
        }
        catch(error) {
            toast.error(`error: ${error.message}`);
        }
        finally {

        }
    };


    const renderForm = () => {
        return (
            <form onSubmit={handleSubmit}>
                <div>
                    <label>
                        Name:
                        <input type="text" value={name} onChange={(e) => setName(e.target.value)} />
                    </label>
                </div>
                <div>
                    <label>
                        Secure:
                        <select value={secure} onChange={(e) => setSecure(e.target.value)}>
                            <option value="detected">Detected</option>
                            <option value="update">Update</option>
                            <option value="secure">Secure</option>
                        </select>
                    </label>
                </div>
                <div>
                    <label>
                        AllowedGenerate:
                        <input
                            type="checkbox"
                            checked={isAllowedGenerate}
                            onChange={(e) => setIsAllowedGenerate(e.target.checked)}
                        />
                    </label>
                </div>
                <button type="submit">Submit</button>
            </form>
        );
    };

    return (
        <div>
    <table>
        <thead>
        <tr>
            <th>Cheat</th>
            <th>Status</th>
            <th>AllowedGenerate</th>
        </tr>
        </thead>
        <tbody>
        {cheats.map((key,index) => (
            <tr key={key.id}>
                <td>{key.name}</td>
                <td>
                    <select
                        value={key.secure}
                        onChange={(e) => handleSecureChange(key.id, e.target.value)}
                    >
                        <option value="secure">Secure</option>
                        <option value="detected">Detected</option>
                        <option value="update">Update</option>
                    </select>
                </td>
                <td>
                    <select
                        value={key.is_allowed_generate}
                        onChange={(e) => handleAllowedGenerateChange(key.id, e.target.value)}
                    >
                        <option value={1}>Allowed</option>
                        <option value={0}>Forbidden</option>
                    </select>

                </td>
            </tr>
        ))}
        </tbody>
    </table>
            <h2>Create cheat</h2>
            {renderForm()}
        </div>

    )
}

export default Cheats;