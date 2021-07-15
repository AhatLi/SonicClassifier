import PropTypes from 'prop-types';
import React from 'react';

// material-ui
import { makeStyles } from '@material-ui/styles';
import { useTheme } from '@material-ui/styles';
import { Avatar, Box, ButtonBase, Card, Typography } from '@material-ui/core';

// assets
import { IconMenu2 } from '@tabler/icons';

const ColorBox = ({ bgcolor, title, data, dark }) => {
    const theme = useTheme();
    return (
        <React.Fragment>
            <Card sx={{ width: '100%' }}>
                <Box
                    sx={{
                        display: 'flex',
                        justifyContent: 'center',
                        alignItems: 'center',
                        py: 0.5,
                        bgcolor: bgcolor,
                        borderRadius: '10px',
                        margin: '5px',
                        color: dark ? theme.palette.grey[800] : '#ffffff'
                    }}
                >
                    {title && (
                        <Typography variant="h1" color="inherit">
                            {title}
                        </Typography>
                    )}
                    {!title && <Box></Box>}
                </Box>
            </Card>
        </React.Fragment>
    );
};

ColorBox.propTypes = {
    bgcolor: PropTypes.string,
    title: PropTypes.string,
    data: PropTypes.object.isRequired,
    dark: PropTypes.bool
};

// style constant
const useStyles = makeStyles((theme) => ({
    grow: {
        flexGrow: 1
    },
    headerAvatar: {
        ...theme.typography.commonAvatar,
        ...theme.typography.mediumAvatar,
        transition: 'all .2s ease-in-out',
        background: theme.palette.secondary.light,
        color: theme.palette.secondary.dark,
        '&:hover': {
            background: theme.palette.secondary.dark,
            color: theme.palette.secondary.light
        }
    },
    boxContainer: {
        display: 'flex',
        [theme.breakpoints.down('md')]: {
        }
    }
}));

//-----------------------|| MAIN NAVBAR / HEADER ||-----------------------//

const Header = ({ handleLeftDrawerToggle }) => {
    const classes = useStyles();
    const theme = useTheme();

    return (
        <React.Fragment>
            {/* logo & toggler button */}
            <div className={classes.boxContainer}>
                <ButtonBase sx={{ borderRadius: '12px', overflow: 'hidden' }}>
                    <Avatar variant="rounded" className={classes.headerAvatar} onClick={handleLeftDrawerToggle} color="inherit">
                        <IconMenu2 stroke={1.5} size="1.3rem" />
                    </Avatar>
                </ButtonBase>
            </div>
            
            <ColorBox 
                bgcolor={theme.palette.primary.light}
                data={{ label: 'Shade-50', color: theme.palette.primary.light }}
                title="SonicClassifier"
                dark
            />
        </React.Fragment>
    );
};

Header.propTypes = {
    handleLeftDrawerToggle: PropTypes.func
};

export default Header;
