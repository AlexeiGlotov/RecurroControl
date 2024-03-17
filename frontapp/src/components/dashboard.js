
import React, {useState, useEffect, useContext} from 'react';
import {Container, Row, Col, Form, Table, Card, Button} from 'react-bootstrap';
import {axiosInstanceWithJWT} from "../api/axios";
import handleError from "../utils/errorHandler";
import {AuthContext} from "./AuthContext";

function Dashboard() {

    const [data, setData] = useState({ count_active: 0, count_all: 0, count_ban: 0 });
    const [date, setDate] = useState(new Date().toISOString().split('T')[0]); // Сегодняшняя дата в формате YYYY-MM-DD
    const [isLoading, setIsLoading] = useState(false); // Состояние загрузки данных
    const {role } = useContext(AuthContext);

    useEffect(() => {
        fetchData();
    }, [date]);

    const fetchData = async () => {
        setIsLoading(true); // Начало загрузки данных
        try {
            const res = await axiosInstanceWithJWT.get(`/api/license-keys/getCustomLicenseKey?date=${date}`);
            setData(res.data.res);
            setIsLoading(false); // Конец загрузки данных
        } catch (error) {
            handleError(error)
            setIsLoading(false);
        }
    };

    const handleBanAll = async () => {
            try {
                await axiosInstanceWithJWT.post('/api/license-keys/ban-of-date', {"date": date});
                await fetchData()
            }
            catch(error) {
                handleError(error);
            }
            finally {
            }
    };

    return (
        <Container className="my-4">
            <Card className="p-3">
            <Row className="mb-3">
                <Col>
                    <Form.Group>
                        <Form.Control
                            type="date"
                            value={date}
                            onChange={(e) => setDate(e.target.value)}
                        />
                    </Form.Group>
                </Col>
            </Row>
            <Row>
                <Col>
                    <Table striped bordered hover>
                        <thead>
                        <tr>
                            <th>Count keys all</th>
                            <th>Keys activated</th>
                            <th>Keys banned</th>
                            {(role === 'admin' || role === 'distributors') && (
                                <th>Action</th>
                            )}
                        </tr>
                        </thead>
                        <tbody>
                        {isLoading ? (
                            <tr>
                                <td colSpan="4">Loading...</td>
                            </tr>
                        ) : data.count_all === 0 ? (
                            <tr>
                                <td colSpan="4">No keys available</td>
                            </tr>
                        ) : (
                            <tr>
                                <td>{data.count_all}</td>
                                <td>{data.count_active}</td>
                                <td>{data.count_ban}</td>
                                <td>
                                    <Button
                                        onClick={() => handleBanAll()}
                                        variant="outline-success"
                                        className="me-2"
                                    >
                                        <i className="bi bi-lock-fill"></i>Banned all
                                    </Button>
                                </td>
                            </tr>
                        )}
                        </tbody>
                    </Table>
                </Col>
            </Row>
            </Card>
        </Container>
    );
}

export default Dashboard;