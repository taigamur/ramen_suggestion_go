import { memo , VFC } from "react";
import {Route, Switch} from "react-router-dom"
import React from "react"

import { Home } from "../components/pages/Home"
import { Page404 } from "../components/pages/Page404";

export const Router = () => {
    return (
        <Switch>
            <Route exact path="/">
                <Home />
            </Route>
            <Route path="*">
                <Page404 />
            </Route>
        </Switch>
    );
};