import React, {useContext, useEffect, useState} from 'react';
import {axiosInstanceWithJWT} from "../api/axios";
import {Table, Button, Form, Card,Container} from 'react-bootstrap';
import { AuthContext } from '../components/AuthContext';
import handleError from '../utils/errorHandler';

// TO:DO Динамическое добавление в таблицу

function AccessKeys() {

    const [sendrole, setRole] = useState('');
    const [isLoading, setIsLoading] = useState(false);
    const [isLoadingGET, setIsLoadingGET] = useState(false);
    const [keys, setKeys] = useState(null);

    const [allKeys, setAllKeys] = useState(null);
    const { role } = useContext(AuthContext);

    const handleRoleChange = (event) => {
        setRole(event.target.value);
    };

    useEffect(() => {
        const fetchUsers = async () => {
            setIsLoadingGET(true)
            try {
                const response = await axiosInstanceWithJWT.get('/api/access-keys/');
                setAllKeys(response.data.key);
            } catch (error) {
                handleError(error);
            } finally {
                setIsLoadingGET(false)
            }
        };

        fetchUsers();
    }, []);

    const handleSubmit = async () => {
        setIsLoading(true);
        const dataToSend = {
            role:sendrole,
        };

        try {
            const [licenseResponse] = await Promise.all([
                await axiosInstanceWithJWT.post('/api/access-keys/',dataToSend),
            ]);

            const response = await axiosInstanceWithJWT.get('/api/access-keys/');
            setAllKeys(response.data.key);

            setKeys(licenseResponse.data.key)
        } catch (error) {
            handleError(error);
        } finally {
            setIsLoading(false);
        }
    };

    return (
        <div>
            <div className="container my-4">
                <Card className="p-3">
                    <div className="d-flex align-items-stretch">
                        <Form.Control type="text" className="me-3" readOnly value={isLoading ? "Loading..." : keys || ''}
                                      style={{maxWidth: '600px'}}/>

                        <Form.Group className="me-3" style={{width: '200px'}}>
                            <Form.Control as="select" id="roleSelect" value={sendrole} onChange={handleRoleChange}>
                                <option value="">select role</option>
                                {role === 'admin' && (
                                    <>
                                        <option value="admin">admin</option>
                                        <option value="distributors">distributors</option>
                                    </>
                                )}
                                {(role === 'distributors' || role === 'admin') && (
                                    <>
                                        <option value="reseller">reseller</option>
                                        <option value="salesman">salesman</option>
                                    </>
                                )}

                                {role === 'reseller' && (
                                    <>
                                        <option value="salesman">salesman</option>
                                    </>
                                )}

                            </Form.Control>
                        </Form.Group>
                        <Button variant="outline-success" onClick={handleSubmit} className="me-3"
                                style={{width: '150px', whiteSpace: 'nowrap'}}>Generate</Button>
                    </div>
                </Card>
            </div>


            <Container className="my-4">
                <Card className="p-3">
                    <Table striped bordered hover> {/* Измененный стиль таблицы */}
                        <thead>
                        <tr>
                            {role === 'admin' && (
                                <th>ID</th>
                            )}
                            <th>Access Key</th>
                            <th>Owner</th>
                            <th>Is Login</th>
                            <th>Role</th>
                        </tr>
                        </thead>
                        <tbody>
                        {allKeys && allKeys.length > 0 ? (
                            allKeys.map((key) => (
                                <tr key={key.id}>
                                    {role === 'admin' && (
                                    <td>{key.id}</td>
                                    )}
                                    <td>{key.access_key}</td>
                                    <td>{key.owner}</td>
                                    <td>{key.is_login}</td>
                                    <td>{key.role}</td>
                                </tr>
                            ))
                        ) : (
                           <>
                               {
                                   isLoadingGET ? (
                                       <tr>
                                           <td colSpan="5">Loading.....</td>
                                       </tr>
                                   ) : (
                                       <tr>
                                           <td colSpan="5">No keys available</td>
                                       </tr>
                                   )
                               }
                           </>
                        )}
                        </tbody>
                    </Table>
                </Card>
            </Container>
        </div>
    )
}

export default AccessKeys;