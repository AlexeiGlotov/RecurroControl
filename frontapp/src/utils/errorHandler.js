// src/utils/errorHandler.js

import { toast } from 'react-toastify';

//handleError(error);
// import handleError from '../utils/errorHandler';
const handleError = (error) => {
    if (error.response && error.response.data && error.response.data.message) {
        switch (error.response.data.message) {
            case "not all fields are filled in":
                toast.error("not all fields are filled in");
                break;
            case "Access denied":
                toast.error("Access denied");
                break;
            case "bad len login":
                toast.error("bad len login 4-20");
                break
            case "bad len password":
                toast.error("bad len password 6-20");
                break
            case "check the repassword != passwords are correct":
                toast.error("check the repassword != passwords are correct");
                break
            case "invalid key":
                toast.error("check access key are correct");
                break
            case "enter another login":
                toast.error("enter another login");
                break
            default:
                toast.error("server error");
        }
    } else {
        toast.error("An unexpected error occurred");
    }
}

export default handleError;