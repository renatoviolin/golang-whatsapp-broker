import React from 'react'
import ChatBody from '../chat-body/ChatBody'
import ChatFooter from '../chat-footer/ChatFooter'
import ChatHeader from '../chat-header/ChatHeader'
import './chatContainer.css'

function ChatContainer() {
   return (
      <div className='chat-container'>
         <ChatHeader />
         <ChatBody/>
         <ChatFooter />
      </div>
   )
}

export default ChatContainer