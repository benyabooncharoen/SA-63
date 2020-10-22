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
    EntSystemmemberEdges,
    EntSystemmemberEdgesFromJSON,
    EntSystemmemberEdgesFromJSONTyped,
    EntSystemmemberEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntSystemmember
 */
export interface EntSystemmember {
    /**
     * Password holds the value of the "Password" field.
     * @type {string}
     * @memberof EntSystemmember
     */
    password?: string;
    /**
     * 
     * @type {EntSystemmemberEdges}
     * @memberof EntSystemmember
     */
    edges?: EntSystemmemberEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntSystemmember
     */
    id?: number;
    /**
     * SystemmemberName holds the value of the "systemmemberName" field.
     * @type {string}
     * @memberof EntSystemmember
     */
    systemmemberName?: string;
}

export function EntSystemmemberFromJSON(json: any): EntSystemmember {
    return EntSystemmemberFromJSONTyped(json, false);
}

export function EntSystemmemberFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntSystemmember {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'password': !exists(json, 'Password') ? undefined : json['Password'],
        'edges': !exists(json, 'edges') ? undefined : EntSystemmemberEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'systemmemberName': !exists(json, 'systemmemberName') ? undefined : json['systemmemberName'],
    };
}

export function EntSystemmemberToJSON(value?: EntSystemmember | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Password': value.password,
        'edges': EntSystemmemberEdgesToJSON(value.edges),
        'id': value.id,
        'systemmemberName': value.systemmemberName,
    };
}

