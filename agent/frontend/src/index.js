import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import Main from './components/main/Main';
import { ChatContextProvider } from './context/chat-context';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
   <ChatContextProvider>
      <Main />
   </ChatContextProvider>
);