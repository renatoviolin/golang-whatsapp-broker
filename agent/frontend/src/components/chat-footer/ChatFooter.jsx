import React, { useContext, useEffect, useRef, useState } from 'react'
import ChatContext from '../../context/chat-context'
import './chatFooter.css'

function ChatFooter() {
   const inputRef = useRef();
   const [text, setText] = useState("")
   const { addAgentMessage } = useContext(ChatContext)

   const handleSubmit = async () => {
      if (text.trim() === "") {
         return
      }
      addAgentMessage(text, "")
      setText("")
   }

   const handleKeyUp = (e) => {
      if (e.code === "Enter") {
         handleSubmit()
      }
   }

   useEffect(() => {
      inputRef.current.focus();
   }, [])

   return (
      <div className='footer-body'>
         <div className="footer-input-container">
            <input
               onChange={(e) => { setText(e.target.value) }}
               onKeyUp={handleKeyUp}
               value={text}
               className='user-input'
               type="text"
               name="user-input"
               id="user-input" ref={inputRef} />
         </div>
         <div className="footer-button-container">
            <button className='btn-submit' name="btn-submit" onClick={handleSubmit}>Send </button>
         </div>
      </div>
   )
}

export default ChatFooter