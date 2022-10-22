import React from 'react'
import './textMessage.css'
import user from './user.png'
import agent from './agent.png'

function TextMessage(props) {
   let render = ""
   if (props.from === "user") {
      render = (
         <div className='message user'>
            <div className="avatar"><img src={user} alt="" /></div>
            <span className='text'>{props.children}</span>
         </div>
      )
   } else if (props.from === "agent") {
      render = (
         <div className='message agent'>
            <span className='text'>{props.children}</span>
            <div className="avatar"><img src={agent} alt="" /></div>
         </div>
      )
   } else if (props.from === "error") {
      render = (
         <div className='message error'>
            <span className='text'>{props.children}</span>
         </div>
      )
   }

   return (
      render
   )
}

export default TextMessage