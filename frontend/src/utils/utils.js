const calculateProgress = (current, goal) => {
    return Math.floor(current / goal * 100);
}

const calculateDayDifference = (dateA, dateB) => {
    return Math.round((dateA - dateB) / (1000 * 60 * 60 * 24));
}

const parseDateFromRFC3339 = (date) => {
    return Date.parse(date)
}

const getDaySince = (date) => {
    return calculateDayDifference(Date.now(), parseDateFromRFC3339(date));
}

export default { calculateProgress, calculateDayDifference, parseDateFromRFC3339, getDaySince }
