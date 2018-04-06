import React, { Component } from "react";
import { connect } from "react-redux";

import { fetchMissions } from "../../actions/missions/actions";

class MissionsProvider extends Component {
    componentDidMount() {
        this.props.fetchMissions();
    }

    render() {
        return this.props.render({ missions: this.props.missions });
    }
}

const mapStateToProps = ({ missions }) => {
    return {
        missions
    };
};

const mapDispatchToProps = {
    fetchMissions
};

export default connect(mapStateToProps, mapDispatchToProps)(MissionsProvider);
