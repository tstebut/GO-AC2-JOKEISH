class App extends React.Component {
  render() {
    console.log("Starting App");
    return <Home />;
  }
}

class Home extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      jokes: []
    };
    this.serverRequest = this.serverRequest.bind(this);
  }

  serverRequest() {
    $.get("/api/jokes", res => {
      console.log("res... ", res);
      this.setState({
        jokes: res
      });
    });
  }

  componentDidMount() {
    this.serverRequest();
  }

  render() {
    return (
      <div className="container">
        <br />
        <h2>Mais c'est une blague ??!</h2>
        <p>C'est parti pour quelques blagues !!!</p>
        <div className="row">
          <div className="joke-container">
            {this.state.jokes.map(function(joke, i) {
              return <Joke key={i} joke={joke} />;
            })}
          </div>
        </div>
      </div>
    );
  }
}

class Joke extends React.Component {
  render() {
    return (
        <div className="panel panel-default joke-ctn">
          <div className="panel-heading">
            #{this.props.joke.id}{" "}
          </div>
          <div className="panel-body joke-hld">{this.props.joke.joke}</div>
          <div className="panel-footer">
            (c) Sentelis
          </div>
        </div>
    );
  }
}
ReactDOM.render(<App />, document.getElementById("app"));
