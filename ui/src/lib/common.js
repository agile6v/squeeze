
import ky from 'ky'

export function getAppRoot() {
    // if (process.env.NODE_ENV !== 'production') {
    //     return 'http://localhost:3000'
    // }
    return ''
}

export const request = ky.extend({ prefixUrl: getAppRoot() + '/api/' })
