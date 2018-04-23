import React, { Component } from "react";
import { connect } from "react-redux";

import { fetchAircraft } from "../../actions/aircraft/actions";

class AircraftProvider extends Component {
    componentDidMount() {
        this.props.fetchAircraft(null, this.props.status);
    }

    render() {
        return this.props.render({
            fetchAircraft: this.props.fetchAircraft,
            aircraft: this.props.aircraft
        });
    }
}

const mapStateToProps = ({ aircraft }) => {
    return {
        aircraft
    };
};

const mapDispatchToProps = {
    fetchAircraft
};

export default connect(mapStateToProps, mapDispatchToProps)(AircraftProvider);
