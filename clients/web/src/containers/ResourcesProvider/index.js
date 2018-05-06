import React, { Component } from "react";
import { connect } from "react-redux";

import { fetchResources } from "../../actions/resources/actions";

class ResourcesProvider extends Component {
    componentDidMount() {
        this.props.fetchResources();
    }

    render() {
        return this.props.render({ resources: this.props.resources });
    }
}

const mapStateToProps = ({ resources }) => {
    return {
        resources
    };
};

const mapDispatchToProps = {
    fetchResources
};

export default connect(mapStateToProps, mapDispatchToProps)(ResourcesProvider);
