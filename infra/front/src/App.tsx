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
      <header className="App-header">
        <BrowserRouter>
          <Banner />
          <Switch>
            <Route exact path="/react">
              <SignInForm />
            </Route>
            <Route exact path="/react/signup">
              <SignupForm />
            </Route>
            <Route exact path="/react/list">
              <FavList />
            </Route>
            <Route exact path="/react/fav">
              <FavForm />
            </Route>
          </Switch>
        </BrowserRouter>
      </header>
    </div>
  );
};

export default App;
