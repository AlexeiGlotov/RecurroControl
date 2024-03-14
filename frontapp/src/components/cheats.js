import React, {useEffect, useState} from "react";
import {axiosInstanceWithJWT} from "../api/axios";
import {toast} from "react-toastify";
import {Button, Card, Container, Form, Table} from "react-bootstrap";

function Cheats() {
    const [cheats, setCheats] = useState([]);
    const [name, setName] = useState('');
    const [secure, setSecure] = useState('secure');
    const [isAllowedGenerate, setIsAllowedGenerate] = useState(1);

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
        const formData = { name, secure, is_allowed_generate: parseInt(isAllowedGenerate) };
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

    return (
    <div>

        <Container className="my-4">
            <Card className="p-3 mb-3">
                <Form className="d-flex align-items-stretch">
                    <Form.Group className="me-3 flex-grow-1">
                        <Form.Control type="text" placeholder="name" value={name} onChange={(e) => setName(e.target.value)} />
                    </Form.Group>
                    <Form.Group className="me-3" style={{ width: '200px' }}>
                        <Form.Control as="select" id="roleSelect" value={secure} onChange={(e) => setSecure(e.target.value)}>
                            <option value="detected">Detected</option>
                            <option value="update">Update</option>
                            <option value="secure">Secure</option>
                        </Form.Control>
                    </Form.Group>
                    <Form.Group className="me-3" style={{ width: '200px' }}>
                        <Form.Control as="select" value={isAllowedGenerate} onChange={(e) => setIsAllowedGenerate(e.target.value)}>
                            <option value="1">Allowed</option>
                            <option value="0">Forbidden</option>
                        </Form.Control>
                    </Form.Group>
                    <Button variant="outline-success" onClick={handleSubmit} className="ms-auto" style={{ width: '150px', whiteSpace: 'nowrap' }}>
                        Generate
                    </Button>
                </Form>
            </Card>

            <Card className="p-3 ">
            <Table striped bordered hover> {/* Измененный стиль таблицы */}
                    <thead>
                    <tr>
                        <th>ID</th>
                        <th>Cheat</th>
                        <th>Status</th>
                        <th>Allowed</th>
                    </tr>
                    </thead>
                    <tbody>
                    {cheats && cheats.length > 0 ? (
                        cheats.map((cheat) => (
                            <tr key={cheat.id}>
                                <td>{cheat.id}</td>
                                <td>{cheat.name}</td>

                                <td>
                                    <select
                                        value={cheat.secure}
                                        onChange={(e) => handleSecureChange(cheat.id, e.target.value)}
                                    >
                                        <option value="secure">Secure</option>
                                        <option value="detected">Detected</option>
                                        <option value="update">Update</option>
                                    </select>
                                </td>
                                <td>
                                    <select
                                        value={cheat.is_allowed_generate}
                                        onChange={(e) => handleAllowedGenerateChange(cheat.id, e.target.value)}
                                    >
                                        <option value={1}>Allowed</option>
                                        <option value={0}>Forbidden</option>
                                    </select>

                                </td>
                            </tr>
                        ))
                    ) : (
                        <tr>
                            <td colSpan="5">No keys available</td>
                        </tr>
                    )}
                    </tbody>
                </Table>
            </Card>
        </Container>
    </div>
    )
}

export default Cheats;