import {Route, Switch} from "react-router-dom"
import { Home } from "../components/pages/Home"
import { Page404 } from "../components/pages/Page404";
import { Index } from "../components/pages/Index";
import { Root } from "../components/pages/Root" 
import { HeaderLayout } from "../components/templates/HeaderLayout";
import { Signup } from "../components/pages/Signup";
import { Login } from "../components/pages/Login"
import { Places } from "../components/pages/Places"
import { UserIndex } from "../components/pages/UserIndex";
import { LoginUserProvider } from "../providers/LoginUserProvider";
import { UserInfo } from "../components/pages/UserInfo"; 

export const Router = () => {
    return (
        <Switch>
            <LoginUserProvider>
                <Route exact path="/"><Root /></Route>
                <Route exact path="/home"><HeaderLayout><Home /></HeaderLayout></Route>
                <Route exact path="/index"><HeaderLayout><Index /></HeaderLayout></Route>
                <Route exact path="/login"><Login/></Route>
                <Route exact path="/signup"><Signup/></Route>
                <Route exact path="/places"><HeaderLayout><Places/></HeaderLayout></Route>
                <Route exact path="/user/index"><UserIndex/></Route>
                <Route exact path="/user/info"><HeaderLayout><UserInfo /></HeaderLayout></Route>
            </LoginUserProvider>
            <Route path="*"><Page404 /></Route>
        </Switch>
    );
};