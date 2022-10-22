import React from 'react'
import avatar from '../text-message/client.png';
import LoadingIcon from './loading.gif'

function Loading() {
   return (
      <div className='message-bot'>
         <div className={`message-bot-avatar`} >
            <img src={avatar} className='bot-avatar' alt='' />
         </div>
         <div className='message-bot-texto align-bottom'>
            <img src={LoadingIcon} alt="" className='loading' />
         </div>
      </div>
   )
}

export default Loading