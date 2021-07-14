import React from 'react';
import { Redirect, Switch } from 'react-router-dom';

// routes
import MainRoutes from './MainRoutes';

// project imports
import config from './../config';

//-----------------------|| ROUTING RENDER ||-----------------------//

const Routes = () => {
    return (
        <Switch>
            <Redirect exact from="/" to={config.defaultPath} />
            <React.Fragment>
                <MainRoutes />
            </React.Fragment>
        </Switch>
    );
};

export default Routes;
