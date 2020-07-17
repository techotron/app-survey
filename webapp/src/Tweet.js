import React from 'react';
import './App.css';

function Tweet(props) {
  return(
    <div className="tweet">
      <h3>{props.name}</h3>
      <p>{props.message}</p>
      <h3>{Likes()}</h3>
    </div>
  )
}

function Likes() {
  return(
    Math.round(Math.random() * 10)
  )
}

export default Tweet;
