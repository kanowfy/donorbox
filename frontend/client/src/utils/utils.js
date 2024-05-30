const formatNumber = (num) => {
    return num.toLocaleString();
}

const calculateProgress = (current, goal) => {
    const progress = Math.floor(current / goal * 100);
    return progress < 100 ? progress : 100;
}

const calculateDayDifference = (dateA, dateB) => {
    return Math.round((dateA - dateB) / (1000 * 60 * 60 * 24));
}

const parseDateFromRFC3339 = (date) => {
    return Date.parse(date)
}

const getRFC3339DateString = (date) => {
    return date.toISOString();
}

const getDaySince = (date) => {
    return calculateDayDifference(Date.now(), parseDateFromRFC3339(date));
}

const formatDate = (date) => {
    let year = date.getFullYear();
    let month = (1 + date.getMonth()).toString().padStart(2, '0');
    let day = date.getDate().toString().padStart(2, '0');

    return day + '/' + month + '/' + year;
}

const parseExpiry = (date) => {
    return {
        month: parseInt(date.substring(0, 2), 10),
        year: parseInt("20" + date.substring(2), 10)
    }
}

const getProvincesByCountry = (data, country) => {
    const record = data.filter(d => d.name === country)[0];
    return record?.states.map(s => s.name);
}

export default { formatNumber, calculateProgress, calculateDayDifference, parseDateFromRFC3339, getRFC3339DateString, getDaySince, formatDate, parseExpiry, getProvincesByCountry }
