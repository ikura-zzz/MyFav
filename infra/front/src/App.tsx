import React from 'react';
import { SignInForm } from './components/SignInForm';
import { Banner } from './components/Banner';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import { SignupForm } from './components/SignupForm';
import { FavList } from './components/FavList';
import { FavForm } from './components/FavForm';

const App = () => {
  return (
    <div className="App">
      <header className="App-header"></header>
      <body className="body">
        <BrowserRouter>
          <Banner />
          <Switch>
            <Route exact path="/react">
              <title>Sign In</title>
              <SignInForm />
            </Route>
            <Route exact path="/react/signup">
              <title>Sign Up</title>
              <SignupForm />
            </Route>
            <Route exact path="/react/list">
              <title>Fav List</title>
              <FavList />
            </Route>
            <Route exact path="/react/fav">
              <title>Fav Form</title>
              <FavForm />
            </Route>
          </Switch>
        </BrowserRouter>
      </body>
    </div>
  );
};

export default App;
