import React, { useContext } from 'react'
import ChatContext from '../../context/chat-context'
import './chatHeader.css'

function ChatHeader() {
   const { connected } = useContext(ChatContext)

   return (
      <div className='header-title'>
         {
            connected === true ? ("Websocket: connected") : ("Websocket: Not connected")
         }
      </div >
   )
}

export default ChatHeader