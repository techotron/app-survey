import React from 'react';
import Tweet from './Tweet';

function App() {

  return (
    <div className="app">
      <Tweet name="Eddy Chorizo" message="Bro, this chorizo is off the hook." />
      <Tweet name="Pete Cheese" message="This cheese needs some goddam butter." />
      <Tweet name="James Pasta" message="Hey, I make more than just pasta!" />
      <Tweet name="Leo Tom Yum" message="More chillis in that please." />
    </div>
  )
}

export default App;
