import React, { Component } from "react";
import { connect } from "react-redux";

import { fetchGroups } from "../../actions/groups/actions";

class GroupsProvider extends Component {
    componentDidMount() {
        this.props.fetchGroups();
    }

    render() {
        return this.props.render({ groups: this.props.groups });
    }
}

const mapStateToProps = ({ groups }) => {
    return {
        groups
    };
};

const mapDispatchToProps = {
    fetchGroups
};

export default connect(mapStateToProps, mapDispatchToProps)(GroupsProvider);
