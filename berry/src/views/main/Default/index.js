import React, { useEffect, useState } from 'react';

// material-ui
import { Grid } from '@material-ui/core';

// project imports
import { gridSpacing } from './../../../store/constant';

const Main = () => {
    const [isLoading, setLoading] = useState(true);
    useEffect(() => {
        setLoading(false);
    }, []);

    return (
        <Grid container spacing={gridSpacing}>
            <Grid item xs={12}>
                <Grid container spacing={gridSpacing}>
                    
                </Grid>
            </Grid>
            <Grid item xs={12}>
            </Grid>
        </Grid>
    );
};

export default Main;
