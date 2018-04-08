import React, { Component } from "react";
import { connect } from "react-redux";

import { fetchMissionDetail } from "../../actions/missionDetail/actions";

class MissionDetailProvider extends Component {
    componentDidMount() {
        this.props.fetchMissionDetail(this.props.id);
    }

    render() {
        return this.props.render({ missionDetail: this.props.missionDetail });
    }
}

const mapStateToProps = ({ missionDetail }) => {
    return {
        missionDetail
    };
};

const mapDispatchToProps = {
    fetchMissionDetail
};

export default connect(mapStateToProps, mapDispatchToProps)(
    MissionDetailProvider
);
