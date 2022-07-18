import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import AppFooter from './AppFooter';
import AppContent from './AppContent';
import AppHeader from './AppHeader';

import 'bootstrap/dist/css/bootstrap.min.css';
import 'bootstrap/dist/js/bootstrap.bundle.min.js';
import './index.css';

class App extends Component {

  render() {
    const myProps = {
      title: "my cool app",
      subject: "my subject",
      favorite_color: "blue"
    }

    return (
      <div className="app">
        <AppHeader {...myProps} />
        <AppContent />
        <AppFooter />
      </div>
    );
  }
}

ReactDOM.render(<App />, document.getElementById('root'));
