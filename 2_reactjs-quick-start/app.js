class Channel extends React.Component {
  onClick() {

  }
  render() {
    const { name } = this.props
    return (
      <div>
        <li onClick={this.onClick.bind(this)}>{name}</li>
      </div>
    )
  }

}

class ChannelList extends React.Component {
  render() {
    const { channels } = this.props
    return(
      <ul>
        {channels.map( channel => {
          return <Channel name={channel.name} key={channel.name}/>
        })}
      </ul>
    )
  }
}

class ChannelForm extends React.Component {
  constructor(props) {
    super(props)
    this.state = {}
  }
  onChange(e) {
    this.setState({channelName:e.target.value})
  }
  onSubmit(e) {
    let { channelName } = this.state
    this.setState({channelName:''})
    this.props.addChannel(channelName)
    e.preventDefault()
  }
  render() {
    return(
      <form onSubmit={this.onSubmit.bind(this)}>
        <input type='text'
          onChange={this.onChange.bind(this)}
          value={this.state.channelName}
        />
      </form>
    )
  }
}

class ChannelSection extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      channels: [
        {name: 'Hardware Support'},
        {name: 'Software Support'}
      ]
    }
  }
  addChannel(name) {
    const { channels } = this.state
    channels.push({ name })
    this.setState({ channels })
  }
  render() {
    return(
      <div>
        <ChannelList channels={this.state.channels} />
        <ChannelForm addChannel={this.addChannel.bind(this)}/>
      </div>
    )
  }
}

ReactDOM.render(<ChannelSection />, document.getElementById('app'))
