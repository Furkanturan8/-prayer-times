const getCurrentTime = () => {
    const now = new Date();
    return {
        hours: now.getHours(),
        minutes: now.getMinutes(),
    };
};

const timeToMinutes = (timeString) => {
    if (!timeString) return 0;
    const [timePart] = timeString.split(' (');
    const [hours, minutes] = timePart.split(':').map(Number);
    return hours * 60 + minutes;
};

const minutesToTime = (minutes) => {
    const hours = Math.floor(minutes / 60);
    const mins = minutes % 60;
    return `${hours} saat ${mins} dakika`;
};

export const calculateCountdown = (timings) => {
    const current = getCurrentTime();
    const currentMinutes = current.hours * 60 + current.minutes;

    if (!timings || Object.keys(timings).length === 0) {
        return { nowPrayer: '', nextPrayer: '', countdown: 'Veri Yüklenmedi' };
    }

    const times = [
        { name: 'İmsak', time: timings.imsak },
        { name: 'Sabah', time: timings.sunrise },
        { name: 'Öğle', time: timings.dhuhr },
        { name: 'İkindi', time: timings.asr },
        { name: 'Akşam', time: timings.maghrib },
        { name: 'Yatsı', time: timings.isha },
    ];

    let nowPrayer = '';
    let nextPrayer = '';
    let countdown = '';

    for (let i = 0; i < times.length; i++) {
        const currentPrayer = times[i];
        const nextPrayerObj = times[i + 1] || null;

        const currentPrayerMinutes = timeToMinutes(currentPrayer.time);
        const nextPrayerMinutes = nextPrayerObj ? timeToMinutes(nextPrayerObj.time) : Infinity;

        // Debugging
        console.log(`Current Prayer: ${currentPrayer.name}, Time: ${currentPrayer.time}, Minutes: ${currentPrayerMinutes}`);
        console.log(`Next Prayer: ${nextPrayerObj ? nextPrayerObj.name : 'None'}, Time: ${nextPrayerObj ? nextPrayerObj.time : 'None'}, Minutes: ${nextPrayerMinutes}`);
        console.log(`Current Minutes: ${currentMinutes}`);

        if (currentMinutes >= currentPrayerMinutes && currentMinutes < nextPrayerMinutes) {
            nowPrayer = currentPrayer.name;
            nextPrayer = nextPrayerObj ? nextPrayerObj.name : times[0].name;

            const nextPrayerTime = nextPrayerObj ? timeToMinutes(nextPrayerObj.time) : timeToMinutes(times[0].time);
            if (nextPrayerTime !== null) {
                let minutesUntilNextPrayer = nextPrayerTime - currentMinutes;
                if (minutesUntilNextPrayer < 0) {
                    minutesUntilNextPrayer += 1440; // Günün tamamı
                }
                countdown = minutesToTime(minutesUntilNextPrayer);
            }
            break;
        }

        if (currentPrayer.name === 'Yatsı' && currentMinutes >= currentPrayerMinutes) {
            const nextDayTimes = [{ name: 'İmsak', time: timings.imsak }];
            const nextDayPrayerTime = timeToMinutes(nextDayTimes[0].time);
            const minutesUntilNextDayPrayer = (1440 - currentMinutes) + nextDayPrayerTime;
            nowPrayer = currentPrayer.name;
            nextPrayer = nextDayTimes[0].name;
            countdown = minutesToTime(minutesUntilNextDayPrayer);
            break;
        }
    }

    // Gece yarısı kontrolü
    if (nowPrayer === '') {
        const nextDayTimes = [{ name: 'İmsak', time: timings.imsak }];
        const nextDayPrayerTime = timeToMinutes(nextDayTimes[0].time);
        const minutesUntilNextDayPrayer =  nextDayPrayerTime - currentMinutes ;
        nowPrayer = 'Yatsı';
        nextPrayer = nextDayTimes[0].name;
        countdown = minutesToTime(minutesUntilNextDayPrayer);
    }

    return { nowPrayer, nextPrayer, countdown };
};
