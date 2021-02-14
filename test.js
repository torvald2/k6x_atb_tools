import dateincrement from 'k6/x/atb/dateincrement';
import numincrement from 'k6/x/atb/numericincrement';

const dincrement =  dateincrement.new("2006-02-02",1)
const nincrement = numincrement.new(1,1)

export default function(){
    console.log(dincrement.get())
    console.log(nincrement.get())
}