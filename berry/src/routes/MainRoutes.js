import React, { lazy } from 'react';
import { Route, Switch, useLocation } from 'react-router-dom';

// project imports
import MainLayout from './../layout/MainLayout';
import Loadable from '../ui-component/Loadable';

// dashboard routing
const MainDefault = Loadable(lazy(() => import('../views/main/Default')));

// utilities routing

//-----------------------|| MAIN ROUTING ||-----------------------//

const MainRoutes = () => {
    const location = useLocation();

    return (
        <Route
            path={[
                '/main/default',
            ]}
        >
            <MainLayout>
                <Switch location={location} key={location.pathname}>
                    <Route path="/main/default" component={MainDefault} />
                </Switch>
            </MainLayout>
        </Route>
    );
};

export default MainRoutes;
