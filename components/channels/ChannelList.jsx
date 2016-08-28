import React, { Component, PropTypes } from 'react'
import Channel from './Channel.jsx'

class ChannelList extends Component {

  render() {
    return (
      <ul>{
        this.props.channels.map( channel => {
          return (<Channel
            key={channel.id}
            channel={channel}
            // setChannel={this.props.setChannel}
            // activeChannel={this.props.activeChannel}
            {...this.props}
          />)
        })
      }</ul>
    )
  }

}

ChannelList.propTypes = {
  channels: PropTypes.array.isRequired,
  setChannel: PropTypes.func.isRequired,
  activeChannel: PropTypes.object.isRequired
}

export default ChannelList
