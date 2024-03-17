
import {axiosInstanceWithJWT} from '../api/axios';
import React, {useState, useEffect, useContext} from 'react';
import 'bootstrap-icons/font/bootstrap-icons.css';
import handleError from '../utils/errorHandler';
import { Button, Form, Table, Pagination,Container,Card } from 'react-bootstrap';
import {AuthContext} from "./AuthContext";


function ListLicenseKeys() {
    const {role } = useContext(AuthContext);
    const [isLoading, setLoading] = useState(true);
    const [licenseKeys, setLicenseKeys] = useState([]);
    const [currentPage, setCurrentPage] = useState(1);
    const [filter, setFilter] = useState('');
    const [serverFilter, setServerFilter] = useState("");

    useEffect(() => {
        const fetchLicenseKeys = async () => {
            setLicenseKeys([])
            try {
                const response = await axiosInstanceWithJWT.get(`api/license-keys/?page=${currentPage}&query=${serverFilter}`)
                setLicenseKeys(response.data.keys);
            } catch (error) {
                handleError(error)

            } finally {
                setLoading(false);
            }
        };

        const timerId = setTimeout(() => {
                fetchLicenseKeys();
        }, 500);

        return () => {
            clearTimeout(timerId);
        };// 500 мс задержки

    }, [currentPage,serverFilter]);


    const handlePrevPage = () => {
        if (currentPage > 1) setCurrentPage(currentPage - 1);
    };

    const handleNextPage = () => {
        setCurrentPage(currentPage + 1);
    };


    const handleFilterChange = (event) => {
        setFilter(event.target.value.toLowerCase());
    };

    const handleServerFilterChange = (event) => {
        setServerFilter(event.target.value.toLowerCase());
    };

    // Функция для проверки вхождения фильтра в любое из полей объекта ключа лицензии
    const filterLicenseKeys = (key) => {
        return Object.values(key).some(value => {
            // Если значение равно null, пропускаем его или обрабатываем по-другому
            if (value === null) {
                return false; // Например, пропустить значение null
            }
            // Используем value.toString() только для ненулевых значений
            return value.toString().toLowerCase().includes(filter);
        });
    };

    const filteredLicenseKeys = licenseKeys?.filter(filterLicenseKeys) || [];

    const handleDeleteUser = async (userId) => {
        try {
            await axiosInstanceWithJWT.post('/api/license-keys/delete', {"id": userId});
            setLicenseKeys(licenseKeys.map(key =>
                key.id === userId ? { ...key, is_deleted: 1 } : key));
        }
        catch(error) {
            handleError(error)
        }
        finally {

        }
    };

    const handleBanUnbanUser = async (userId, isBanned) => {
        if (isBanned === 0) {
            try {
                await axiosInstanceWithJWT.post('/api/license-keys/unban', {"id": userId});
                setLicenseKeys(licenseKeys.map(key =>
                    key.id === userId ? { ...key, banned: 0 } : key));
            }
            catch(error) {
                handleError(error)
            }
            finally {

            }
        } else {
            try {
                await axiosInstanceWithJWT.post('/api/license-keys/ban', {"id": userId});
                setLicenseKeys(licenseKeys.map(key =>
                    key.id === userId ? { ...key, banned: 1 } : key));
            }
            catch(error) {
                handleError(error)
            }
            finally {
            }
        }
    };

    const handleResetHWID = async (userId) => {
        try {
            axiosInstanceWithJWT.post('/api/license-keys/resetHWID', {"id": userId});
            setLicenseKeys(licenseKeys.map(key =>
                key.id === userId ? { ...key, hwid: null,hwidk:null } : key));
        }
        catch(error) {
            handleError(error)
        }
        finally {

        }
    };

    const cellStyle = {
        maxWidth: '100px',
        overflow: 'hidden',
        whiteSpace: 'nowrap'
    };

    const renderInfoKeyList = () => {
        return (
            <Container className="my-4">
                <Card className="p-3 mb-3">
                    <div className="d-flex align-items-center mb-3">

                        <Form.Control
                            className="me-3"
                            type="text"
                            value={filter}
                            onChange={handleFilterChange}
                            placeholder="client filter | Example : any word"
                        />

                        <Form.Control
                            className="me-3"
                            type="text"
                            value={serverFilter}
                            onChange={handleServerFilterChange}
                            placeholder="server filter | Example : sql syntax - banned = 1"
                        />

                        <Pagination>
                            <Pagination.Prev onClick={handlePrevPage} disabled={currentPage === 1}/>
                            <Pagination.Item active>{currentPage}</Pagination.Item>
                            <Pagination.Next onClick={handleNextPage}/>
                        </Pagination>
                    </div>


                        <Table responsive striped bordered hover>
                            <thead>
                            <tr>
                                <th>ID</th>
                                <th>License Key</th>
                                <th>Cheat</th>
                                <th>TTL Cheat</th>
                                <th>Holder</th>
                                <th>Creator</th>
                                <th>Date of Creation</th>
                                <th>Date of Activation</th>
                                <th>HWID</th>
                                <th>HWIDK</th>
                                <th>notes</th>
                                <th>ban</th>
                                <th>Action</th>

                            </tr>
                            </thead>
                            <tbody>
                            {licenseKeys && licenseKeys.length > 0 ? (
                            filteredLicenseKeys.filter(key => key.is_deleted !== 1).map(key => (
                                <tr key={key.id}>
                                    <td>{key.id}</td>
                                    <td style={cellStyle}>{key.license_key}</td>
                                    <td>{key.cheat}</td>
                                    <td>{key.ttl_cheat}</td>
                                    <td>{key.holder}</td>
                                    <td>{key.creator}</td>
                                    <td>{key.date_creation}</td>
                                    <td>{key.date_activation}</td>
                                    <td style={cellStyle}>{key.hwid}</td>
                                    <td style={cellStyle}>{key.hwidk}</td>
                                    <td style={cellStyle}>{key.notes}</td>
                                    <td>{key.banned}</td>
                                    <td>
                                        <Button onClick={() => handleResetHWID(key.id, 0)}
                                                disabled={key.hwid === null && key.hwidk === null}
                                                variant="outline-primary"
                                                className="me-2"><i
                                            className="bi bi-server"></i></Button>

                                        {(role === 'admin' || role === 'reseller' || role === 'distributors') && (
                                            <>
                                                {key.banned === 1 ? (
                                                    <Button
                                                        onClick={() => handleBanUnbanUser(key.id, 0)}
                                                        variant="outline-danger"
                                                        className="me-2"
                                                    >
                                                        <i className="bi bi-lock-fill"></i>
                                                    </Button>
                                                ) : (
                                                    <Button
                                                        onClick={() => handleBanUnbanUser(key.id, 1)}
                                                        variant="outline-success"
                                                        className="me-2"
                                                    >
                                                        <i className="bi bi-unlock-fill"></i>
                                                    </Button>
                                                )}
                                            </>
                                        )}

                                        {role === 'admin' && (
                                            <Button onClick={() => handleDeleteUser(key.id)} variant="outline-danger"><i
                                                className="bi bi-trash"></i></Button>
                                        )}
                                    </td>
                                </tr>
                            ))
                            ) : (
                                <>
                                    {
                                        isLoading ? (
                                            <tr>
                                                <td colSpan="12">Loading.....</td>
                                            </tr>
                                        ) : (
                                            <tr>
                                                <td colSpan="12">No keys available</td>
                                            </tr>
                                        )
                                    }
                                </>
                            )}
                            </tbody>
                        </Table>
                </Card>
            </Container>
    )
    };

    return (
        <div>
            {renderInfoKeyList()}
        </div>
    );
}

export default ListLicenseKeys;