import React from 'react'
import {Nav} from 'rsuite';
import Icon from "rsuite/es/Icon";


class HomepageLayout extends React.Component {
    constructor(props) {
        super(props);

        this.setState({kek: "lol"})
    }


    render() {
        return (
            <Nav>
                <Nav.Item icon={<Icon icon="home"/>}>Home</Nav.Item>
                <Nav.Item>News</Nav.Item>
                <Nav.Item>Solutions</Nav.Item>
                <Nav.Item>Products</Nav.Item>
                <Nav.Item>About</Nav.Item>
            </Nav>
        )
    }

}


export default HomepageLayout