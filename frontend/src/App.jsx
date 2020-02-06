import React from 'react';
import axios from 'axios';
import HomepageLayout from "./components/HomePage";


class App extends React.Component {
    constructor(props) {
        super(props);

        this.state = {
            users: []
        }
    }

    componentDidMount() {
        this.getUsers()
    }

    getUsers() {
        axios.get("http://127.0.0.1:8080/user").then((r) => {
            this.setState({users: r.data})
        });
    }

    render() {
        return (
            <HomepageLayout/>
        );
    }
}


export default App;
