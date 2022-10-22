import React, { useContext, useEffect, useRef } from 'react'
import ChatContext from '../../context/chat-context';
import './chatBody.css'

function ChatBody() {
   const { elements, setElements } = useContext(ChatContext)
   const bottomRef = useRef(null);

   useEffect(() => {
      setElements([])
   }, [])

   useEffect(() => {
      bottomRef.current?.scrollIntoView({ behavior: 'smooth' });
   }, [elements]);

   return (
      <div className='chat-body'>
         {
            elements.map((el, i) => {
               return (<div style={{ width: '100%', display: 'flex' }} key={i}> {el} </div>)
            })
         }
         <div ref={bottomRef} className="bottomRef" />
      </div>
   )
}

export default ChatBody