import React, { Component } from "react";
import { connect } from "react-redux";

import { fetchGroupsDetail } from "../../actions/groupsDetail/actions";
import { map } from "rxjs/operators";

class GroupsDetailProvider extends Component {
    componentDidMount() {
        this.props.fetchGroupsDetail(this.props.id);
    }

    render() {
        return this.props.render({ groupsDetail: this.props.groupsDetail });
    }
}

const mapStateToProps = ({ groupsDetail }) => {
    return {
        groupsDetail
    };
};

const mapDispatchToProps = {
    fetchGroupsDetail
};

export default connect(mapStateToProps, mapDispatchToProps)(
    GroupsDetailProvider
);
