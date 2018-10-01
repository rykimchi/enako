import React, { Component } from 'react';
import BigCalendar from 'react-big-calendar';
import moment from 'moment';
import style from 'react-big-calendar/lib/css/react-big-calendar.css';
import '../css/Calendar.css';

const localizer = BigCalendar.momentLocalizer(moment);
const views = ['month'];

class Calendar extends Component {
    onDateSelected = (evt) => {
        const date = moment(evt.start).format('MMMM Do');
        this.props.setSelectedDate(date);
    };

    render() {
        return (
            <div className='Calendar'>
                <BigCalendar
                    localizer={localizer}
                    events={[]}
                    views={views}
                    style={style}
                    selectable='ignoreEvents'
                    onSelectSlot={this.onDateSelected}
                />
            </div>
        );
    }
}

export default Calendar;
