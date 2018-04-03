import React from 'react';
import {Route} from 'react-router-dom'
import Sheets from './sheets-controller.jsx'
import Todo from './todo-controller.jsx'
import link from './dammy/test-sheets.json'


export default class App extends React.Component{

  // link の ajaxを入れる

  render() {
    return (
      <div>
        <nav>
          <Sheets items={link} />
        </nav>
        <section>
          <Route exact path="/todo/" component={Home} />
          <Route path="/todo/:id" component={Todo} />
        </section>
      </div>
    )
  }
}

const Home = () => (
  <h1>HOME</h1>
)

//<Route path="/sheet/:id" component={Sheets} />  

/*
const Todos = props =>{
  console.log(props.match.params)
  return(
  <h1>HOME</h1>
)}
*/