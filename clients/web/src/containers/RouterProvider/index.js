// TODO: Verify that it is indeed good style to use push
// with mapDispatchToProps

import React, { Component } from "react";
import { connect } from "react-redux";

class RouterProvider extends Component {
    render() {
        return this.props.render({ router: this.props.router });
    }
}

const mapStateToProps = ({ router }) => {
    return {
        router
    };
};

export default connect(mapStateToProps)(RouterProvider);
