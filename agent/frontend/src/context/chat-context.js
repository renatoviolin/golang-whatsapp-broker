import React, { createContext, useEffect, useState } from 'react';
import TextMessage from '../components/text-message/TextMessage';
import { w3cwebsocket as W3CWebSocket } from "websocket";

const wsClient = new W3CWebSocket('ws://localhost:7000/websocket');
const ChatContext = createContext(null);


export const ChatContextProvider = ({ children }) => {
   const [connected, setConnected] = useState(false)
   const [elements, setElements] = useState([])

   wsClient.onopen = (msg) => {
      console.log('WebSocket Client Connected');
      setConnected(true)
   };

   wsClient.onclose = () => {
      console.log('WebSocket Client Closed');
      setConnected(false)
   };

   wsClient.onmessage = (message) => {
      const msgJson = JSON.parse(message.data)
      if (msgJson.type === "error") {
         console.error(msgJson)
         addErrorMessage("you are not connected with user")
         return
      }
      addUserMessage(msgJson.body)
   };

   wsClient.onerror = (err) => {
      console.error(err)
   }


   const addAgentMessage = async (text) => {
      wsClient.send(`
            {
               "ws_action": "send",
               "type": "text",
               "text": "${text}"
         }
      `)
      const newElement = <TextMessage from="agent">{text}</TextMessage>
      setElements(oldArray => [...oldArray, newElement]);
   }

   const addErrorMessage = (text) => {
      const newElement = <TextMessage from="error">{text}</TextMessage>
      setElements(oldArray => [...oldArray, newElement]);
   }

   const addUserMessage = (text) => {
      const newElement = <TextMessage from="user">{text}</TextMessage>
      setElements(oldArray => [...oldArray, newElement]);
   }

   useEffect(() => {
      setInterval(() => {
         wsClient.send(`
            {
               "ws_action": "read"
            }
         `)
      }, 1000);
   }, [])

   const value = {
      addAgentMessage,
      elements, setElements,
      connected
   }

   return (
      <ChatContext.Provider value={value}>{children}</ChatContext.Provider>
   )
}

export default ChatContext