import PublicRoutes from "./public";
import React,{useContext} from "react";
import { AuthContext } from '../components/AuthContext';
import PrivateRoutes from "./private";

function Main() {
    const {isAuthenticated} = useContext(AuthContext);

    if (!isAuthenticated){
        return <PublicRoutes />;
    }else{
        return <PrivateRoutes/>;
    }
}

export default Main;