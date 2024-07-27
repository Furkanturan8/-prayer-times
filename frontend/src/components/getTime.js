import {ref} from "vue";

export const time = ref('');
export const secondAngle = ref(0);
export const minuteAngle = ref(0);
export const hourAngle = ref(0);

export const updateTime = () => {
    const now = new Date();
    time.value = now.toLocaleTimeString();
    secondAngle.value = (now.getSeconds() / 60) * 2 * Math.PI;
    minuteAngle.value = ((now.getMinutes() + now.getSeconds() / 60) / 60) * 2 * Math.PI;
    hourAngle.value = ((now.getHours() % 12 + now.getMinutes() / 60) / 12) * 2 * Math.PI;

    return {
        time: time.value,
        secondAngle: secondAngle.value,
        minuteAngle: minuteAngle.value,
        hourAngle: hourAngle.value
    };
};