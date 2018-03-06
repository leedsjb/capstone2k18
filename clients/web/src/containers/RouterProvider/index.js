// TODO: Verify that it is indeed good style to use push
// with mapDispatchToProps

import React, { Component } from "react";
import { connect } from "react-redux";
import { push } from "react-router-redux";

class RouterProvider extends Component {
    render() {
        return this.props.render({ push: this.props.push });
    }
}

const mapDispatchToProps = {
    push
}

export default connect(null, mapDispatchToProps)(RouterProvider);