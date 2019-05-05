const AUTH0_CLIENT_ID = "DNrBHs1rf6DHxfSpxZ0dH8y3fVcevShP";
const AUTH0_DOMAIN = "go-helloworld-go.eu.auth0.com";
const AUTH0_CALLBACK_URL = location.href;
//const AUTH0_API_AUDIENCE = "golang-gin";
const AUTH0_API_AUDIENCE = "https://8081-dot-7170711-dot-devshell.appspot.com/api";

class App extends React.Component {
  parseHash() {
    this.auth0 = new auth0.WebAuth({
      domain: AUTH0_DOMAIN,
      clientID: AUTH0_CLIENT_ID
    });
    this.auth0.parseHash(window.location.hash, (err, authResult) => {
      if (err) {
        return console.log(err);
      }
      if (
        authResult !== null &&
        authResult.accessToken !== null &&
        authResult.idToken !== null
      ) {
        localStorage.setItem("access_token", authResult.accessToken);
        localStorage.setItem("id_token", authResult.idToken);
        localStorage.setItem(
          "profile",
          JSON.stringify(authResult.idTokenPayload)
        );
        window.location = window.location.href.substr(
          0,
          window.location.href.indexOf("#")
        );
      }
    });
  }

  setup() {
    $.ajaxSetup({
      beforeSend: (r) => {
        if (localStorage.getItem("access_token")) {
          r.setRequestHeader(
            "Authorization",
            "Bearer " + localStorage.getItem("access_token")
          );
        }
      }
    });
  }

  setState() {
    let idToken = localStorage.getItem("id_token");
    if (idToken) {
      this.loggedIn = true;
    } else {
      this.loggedIn = false;
    }
  }

  componentWillMount() {
    this.setup();
    this.parseHash();
    this.setState();
  }

  render() {
    if (this.loggedIn) {
      return <LoggedIn />;
    }
    return <Home />;
  }
}

class Home extends React.Component {
  constructor(props) {
    super(props);
    this.authenticate = this.authenticate.bind(this);
  }
  authenticate() {
    this.WebAuth = new auth0.WebAuth({
      domain: AUTH0_DOMAIN,
      clientID: AUTH0_CLIENT_ID,
      scope: "openid profile",
      audience: AUTH0_API_AUDIENCE,
      responseType: "token id_token",
      redirectUri: AUTH0_CALLBACK_URL
    });
    this.WebAuth.authorize();
  }

  render() {
    return (
      <div className="container">
        <div className="row">
          <div className="col-xs-8 col-xs-offset-2 jumbotron text-center">
            <h1>Jokeish</h1>
            <p>Faites le plein de blagues XD</p>
            <p>Connectez-vous</p>
            <a
              onClick={this.authenticate}
              className="btn btn-primary btn-lg btn-login btn-block"
            >
              Connection
            </a>
          </div>
        </div>
      </div>
    );
  }
}

class LoggedIn extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      jokes: []
    };

    this.serverRequest = this.serverRequest.bind(this);
    this.logout = this.logout.bind(this);
  }

  logout() {
    localStorage.removeItem("id_token");
    localStorage.removeItem("access_token");
    localStorage.removeItem("profile");
    location.reload();
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
    console.log("Mounting component");
    this.serverRequest();
    console.log("Ccomponent did mount");
  }

  render() {
    return (
      <div className="container">
        <br />
        <span className="pull-right">
          <a onClick={this.logout}>Se déconnecter</a>
        </span>
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
  constructor(props) {
    super(props);
    this.state = {
      liked: "",
      jokes: []
    };
    this.like = this.like.bind(this);
    this.serverRequest = this.serverRequest.bind(this);
  }

  like() {
    let joke = this.props.joke;
    this.serverRequest(joke);
  }
  serverRequest(joke) {
    $.post(
      "/api/jokes/like/" + joke.id,
      { like: 1 },
      res => {
        console.log("res... ", res);
        this.setState({ liked: "Liked!", jokes: res });
        this.props.jokes = res;
      }
    );
  }

  render() {
    return (
        <div className="panel panel-default joke-ctn">
          <div className="panel-heading">
            #{this.props.joke.id}{" "}
            <span className="pull-right">{this.state.liked}</span>
          </div>
          <div className="panel-body joke-hld">{this.props.joke.joke}</div>
          <div className="panel-footer">
            {this.props.joke.likes} Likes &nbsp;
            <a onClick={this.like} className="btn btn-default">
              <span className="glyphicon glyphicon-thumbs-up" />
            </a>
          </div>
        </div>
    );
  }
}
ReactDOM.render(<App />, document.getElementById("app"));
