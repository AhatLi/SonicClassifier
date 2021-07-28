import { React, useEffect, useState } from 'react';

import MainCard from './../../../ui-component/cards/MainCard';
import { makeStyles } from '@material-ui/styles';
import { FormControl, FormLabel, RadioGroup, FormControlLabel, Radio, Select, MenuItem, InputLabel, Button, Grid } from '@material-ui/core';

const Main = () => {
  const axios = require('axios');

  const [type, setType] = useState('playlist');
  const [item, setItem] = useState('path');
  const [playlist, setPlaylist] = useState("Select Playlist");
  const [playlists, setPlaylists] = useState([]);
  const [order, setOrder] = useState("asc");
  
  const handleChangeType = (event) => {
    setType(event.target.value);
  };
  const handleChangePlaylist = (event) => {
    setPlaylist(event.target.value);
  };
  const handleChangeSortItem = (event) => {
    setItem(event.target.value);
  };
  const handleChangeOrder = (event) => {
    setOrder(event.target.value);
  };

  const useStyles = makeStyles((theme) => ({
    formControl: {
      margin: theme.spacing(1),
      minWidth: 120,
    },
    selectEmpty: {
      marginTop: theme.spacing(2),
    },
    paper: {
      padding: theme.spacing(2),
      textAlign: 'center',
      color: 'blue',
      backgroundColor: 'red'
    },
    frame: {
        height: 'calc(100vh - 210px)',
        border: '1px solid',
        borderColor: theme.palette.primary.light
    },
    area: {
        height: 'calc(80vh - 210px)',
        borderColor: theme.palette.primary.light
    },
    formstyle: {
        textAlign: 'center',
        width: '100%'
    }
  }));
  const classes = useStyles();
    
  const loadingPlaylist = () => {
    axios.get(`/getPlaylist`).then(
      (response) => {        
        const words = response.data.split('|');
        setPlaylists(words);
      }
    ).catch(error => {
        throw(error);
    })
  };
  useEffect(() => {
    loadingPlaylist();
  }, []);

  const onClickSortSong = () => { 
    if({type}.type === "star") {
      axios.get(`/sortStar?item=${item}&order=${order}`).then(
        (response) => {        
          alert(response.data);
        }
      ).catch(error => {
        alert(error);
      });
    }
    else {
      axios.get(`/sortPlaylist?playlist=${playlist}&item=${item}&order=${order}`).then(
        (response) => {        
          alert(response.data);
        }
      ).catch(error => {
        alert(error);
      });
    }

  }

  return (
    <MainCard  className={classes.frame} title="Sort Your Music">
    <Grid container spacing={3} className={classes.area}>
      <Grid item xs={12}>
      <FormControl>
      <FormLabel component="legend">Type</FormLabel>
      <RadioGroup row aria-label="position" name="position" defaultValue="top" value={type} onChange={handleChangeType}>
        <FormControlLabel value="playlist" control={<Radio selected/>} label="Playlist" />
        <FormControlLabel value="star" control={<Radio />} label="Star" />
      </RadioGroup>
      </FormControl>
      </Grid>
      <Grid item xs={12}>
      <InputLabel id="demo-simple-select-helper-label1">Playlist</InputLabel>
      <FormControl className={classes.formstyle}>
        <Select
        id="demo-simple-select1"
        labelId="demo-simple-select-helper-label1"
        value={playlist}
        onChange={handleChangePlaylist}
        inputProps={{ 'aria-label': 'Without label' }}
        >
        <MenuItem value="Select Playlist" disabled>Select Playlist</MenuItem>
        {playlists.map((item, i) => (
          <MenuItem key={i} value={item}>
            {item}
          </MenuItem>
        ))}
        </Select>
      </FormControl>
      </Grid>
      <Grid item xs={12}>
      <InputLabel id="demo-simple-select-helper-label2">Sort Item</InputLabel>
      <FormControl className={classes.formstyle}>
          <Select
          id="demo-simple-select2"
          labelId="demo-simple-select-helper-label2"
          value={item}
          onChange={handleChangeSortItem}
          inputProps={{ 'aria-label': 'Without label' }}
          >
          <MenuItem value='path'>Path</MenuItem>
          <MenuItem value='title'>Title</MenuItem>
          <MenuItem value='album'>Album</MenuItem>
          <MenuItem value='artist'>Artist</MenuItem>
          <MenuItem value='year'>Year</MenuItem>
          </Select>
      </FormControl>
      </Grid>
      <Grid item xs={12}>
      <FormControl>
          <FormLabel component="legend">Order</FormLabel>
          <RadioGroup row aria-label="position" name="position" defaultValue="top" value={order} onChange={handleChangeOrder}>
              <FormControlLabel value="asc" control={<Radio selected/>} label="Asc" />
              <FormControlLabel value="desc" control={<Radio />} label="Desc" />
          </RadioGroup>
      </FormControl>
      </Grid>
      <Grid item xs={12}>
      <Button variant="contained" className={classes.formstyle} onClick={onClickSortSong}>Sort Start!</Button> 
      </Grid>
    </Grid>
    <hr></hr>
    <p style={{textAlign: 'right'}}>
    v0.9.1
    </p>
    </MainCard>
  );
};

export default Main;
