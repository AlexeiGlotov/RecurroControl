import PublicRoutes from "./public";
import React, {useContext, useEffect} from "react";
import { AuthContext } from '../components/AuthContext';
import PrivateRoutes from "./private";
import authEvent from "../components/authEvent";


function Main() {
    const {isAuthenticated,logout} = useContext(AuthContext);

    useEffect(() => {
        const handleUnauthorized = () => {
            logout();
        };

        authEvent.on('unauthorized', handleUnauthorized);
        return () => {
            authEvent.off('unauthorized', handleUnauthorized);
        };
    }, [logout]);

    if (!isAuthenticated){
        return <PublicRoutes />;
    }else{
        return <PrivateRoutes/>;
    }
}

export default Main;