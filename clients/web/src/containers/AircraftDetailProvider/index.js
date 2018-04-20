import React, { Component } from "react";
import { connect } from "react-redux";

import { fetchAircraftDetail } from "../../actions/aircraftDetail/actions";

class AircraftDetailProvider extends Component {
    componentDidMount() {
        this.props.fetchAircraftDetail(this.props.id);
    }

    render() {
        return this.props.render({ aircraftDetail: this.props.aircraftDetail });
    }
}

const mapStateToProps = ({ aircraftDetail }) => {
    return {
        aircraftDetail
    };
};

const mapDispatchToProps = {
    fetchAircraftDetail
};

export default connect(mapStateToProps, mapDispatchToProps)(
    AircraftDetailProvider
);
