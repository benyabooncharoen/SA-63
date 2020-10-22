/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntPatient,
    EntPatientFromJSON,
    EntPatientFromJSONTyped,
    EntPatientToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRightoftreatmentEdges
 */
export interface EntRightoftreatmentEdges {
    /**
     * Patients holds the value of the patients edge.
     * @type {Array<EntPatient>}
     * @memberof EntRightoftreatmentEdges
     */
    patients?: Array<EntPatient>;
}

export function EntRightoftreatmentEdgesFromJSON(json: any): EntRightoftreatmentEdges {
    return EntRightoftreatmentEdgesFromJSONTyped(json, false);
}

export function EntRightoftreatmentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRightoftreatmentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'patients': !exists(json, 'patients') ? undefined : ((json['patients'] as Array<any>).map(EntPatientFromJSON)),
    };
}

export function EntRightoftreatmentEdgesToJSON(value?: EntRightoftreatmentEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'patients': value.patients === undefined ? undefined : ((value.patients as Array<any>).map(EntPatientToJSON)),
    };
}


