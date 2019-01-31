

import React from "react"
import ReactDOM from "react-dom"

import CssBaseline from '@material-ui/core/CssBaseline';
import Drawer from '@material-ui/core/Drawer';
import Toolbar from '@material-ui/core/Toolbar';
import Typography from '@material-ui/core/Typography';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import Hidden from '@material-ui/core/Hidden';

import IconButton from '@material-ui/core/IconButton';
import MenuIcon from '@material-ui/icons/Menu';


import AppBar from '@material-ui/core/AppBar';
import { withStyles } from '@material-ui/core/styles';
import { createMuiTheme } from "@material-ui/core/styles";
import { ThemeProvider } from "@material-ui/styles";

const drawerWidth = 150;

const styles = (theme:any) => ({
  root: {
    display: 'flex',
  },
  drawer: {
    [theme.breakpoints.up('sm')]: {
      width: drawerWidth,
      flexShrink: 0,
    },
  },
  appBar: {
    marginLeft: drawerWidth,
    [theme.breakpoints.up('sm')]: {
      width: `calc(100% - ${drawerWidth}px)`,
    },
  },
  menuButton: {
    marginRight: 20,
    [theme.breakpoints.up('sm')]: {
      display: 'none',
    },
  },
  toolbar: theme.mixins.toolbar,
  drawerPaper: {
    width: drawerWidth,
  },
  content: {
    flexGrow: 1,
    paddingTop: theme.spacing.unit * 2,
    paddingLeft: theme.spacing.unit * 2,
  },
});

function Main({classes, container}: any) {
    return <div className={classes.root}>
        <CssBaseline />
        <AppBar position="fixed" className={classes.appBar}>
            <Toolbar>
              <IconButton
                color="inherit"
                aria-label="Open drawer"
                onClick={() => {}}
                className={classes.menuButton}
              >
                <MenuIcon />
              </IconButton>
              <Typography variant="h6" color="inherit" noWrap>
                Responsive drawer
              </Typography>
            </Toolbar>
        </AppBar>
        <nav className={classes.drawer}>
          <Hidden smUp implementation="css">
              <Drawer
                variant="temporary"
                open={false}
                container={container}
                classes={{
                  paper: classes.drawerPaper,
                }}
              >
                <div>test</div>
              </Drawer>
            </Hidden>
          <Hidden xsDown implementation="css">
              <Drawer open            
                  anchor="left"
                  className={classes.drawer}
                  variant="permanent"
                  classes={{
                    paper: classes.drawerPaper,
                  }}>
                  <List>
                      {['All mail', 'Trash', 'Spam'].map((text) => (
                          <ListItem button key={text}>
                              <ListItemText primary={text} />
                          </ListItem>
                      ))}
                  </List>
              </Drawer>
          </Hidden>
        </nav>
        <main className = {classes.content}>
          <div className={classes.toolbar} />
          <div>Test</div>
          <div>Test</div>
          <div>Test</div>
          <div>Test</div>
          <div>Test</div>
        </main>
    </div>
}
const App = withStyles(styles, {withTheme: true})(Main);

const theme = createMuiTheme({ typography: { useNextVariants: true } });
ReactDOM.render(<ThemeProvider theme={theme}>
    <App />
  </ThemeProvider>, document.getElementById("app"))