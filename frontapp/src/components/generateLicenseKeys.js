import React, {useEffect, useState} from "react";
import {axiosInstanceWithJWT} from "../api/axios";
import {toast} from "react-toastify";
import { Button, Form,Col,Row,Container,Card,InputGroup } from 'react-bootstrap';

function GenerateLicenseKeys() {

    const [keys, setKeys] = useState([]);
    const [isLoading, setIsLoading] = useState(false);
    const [cheats, setCheats] = useState([]);
    const [users, setUsers] = useState([]);
    const [subsCount, setSubsCount] = useState(1);
    const [daysCount, setDaysCount] = useState(25);
    const [selectedCheat, setSelectedCheat] = useState('');
    const [selectedUser, setSelectedUser] = useState('');


    const handleSubmit = async () => {

        setIsLoading(true);
        const dataToSend = {
            count_keys: parseInt(subsCount),
            ttl_cheat: parseInt(daysCount),
            cheat:selectedCheat,
            holder:selectedUser
        };


        try {
            const [licenseResponse] = await Promise.all([
                axiosInstanceWithJWT.post('/api/license-keys/',dataToSend),
            ]);
            setKeys(licenseResponse.data.keys)
        } catch (error) {
            toast.error(`error: ${error.message}`);
        } finally {
            setIsLoading(false);
        }

    };

    useEffect(() => {
        const fetchDatas = async () => {
            try {
                const [cheatsResponse, usersResponse] = await Promise.all([
                    axiosInstanceWithJWT.get('/api/cheats/'),
                    axiosInstanceWithJWT.get('/api/users/getUsers')
                ]);
                setCheats(cheatsResponse.data.cheats);
                setUsers(usersResponse.data.users); // Убедитесь, что здесь правильно обрабатывается ответ
            } catch (error) {
                toast.error(`error: ${error.message}`);
            } finally {
                setIsLoading(false);
            }
        };
        fetchDatas();
    }, []);


    const renderGenerateCheat = () => {
        return (
            <Container className="my-4">
                <Card>
                    <Card.Body>
                        <Row>
                            <Col>
                                <Form.Group>
                                    <Form.Control as="select" value={selectedCheat} onChange={(e) => setSelectedCheat(e.target.value)}>
                                        <option value="">select cheat</option>
                                        {cheats.map((cheat) => (
                                            <option key={cheat.id} value={cheat.name}>
                                                {cheat.name}
                                            </option>
                                        ))}
                                    </Form.Control>
                                </Form.Group>
                            </Col>

                            <Col>
                                <Form.Group>
                                    <Form.Control as="select" value={selectedUser} onChange={(e) => setSelectedUser(e.target.value)}>
                                        <option value="">select user</option>
                                        {users.map((user) => (
                                            <option key={user.id} value={user.login}>
                                                {user.login} ({user.role})
                                            </option>
                                        ))}
                                    </Form.Control>
                                </Form.Group>
                            </Col>

                            <Col>
                                <InputGroup>
                                    <Form.Control type="range" min="0" max="50" value={subsCount} onChange={(e) => setSubsCount(e.target.value)} />
                                    <InputGroup.Text>{subsCount} subs</InputGroup.Text>
                                </InputGroup>
                            </Col>

                            <Col>
                                <InputGroup>
                                    <Form.Control type="range" min="0" max="30" value={daysCount} onChange={(e) => setDaysCount(e.target.value)} />
                                    <InputGroup.Text>{daysCount} day</InputGroup.Text>
                                </InputGroup>
                            </Col>

                            <Col>
                                <Button variant="primary" onClick={handleSubmit}>Generate Subscription</Button>
                            </Col>
                        </Row>
                    </Card.Body>
                </Card>

                <Row className="mt-3">
                    <Col>
                        <Card>
                            <Card.Body>
                                <Card.Title>Keys List</Card.Title>
                                {isLoading ? <p>Loading...</p> : (
                                    keys.map((key, index) => (
                                        <div key={index}>
                                            <p>{key.license_key}</p>
                                        </div>
                                    ))
                                )}
                            </Card.Body>
                        </Card>
                    </Col>
                </Row>
            </Container>
        );
    };

                    return (
        <div>
            {renderGenerateCheat()}

        </div>
    )
}

export default GenerateLicenseKeys;