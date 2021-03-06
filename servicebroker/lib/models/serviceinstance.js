'use strict';
var ServiceInstance = function(sequelize, DataTypes) {
   return sequelize.define('service_instance', {
   serviceId: {
      type: DataTypes.STRING,
      primaryKey: true,
      field: 'service_id',
      allowNull: false
    },
    orgId: {
      type: DataTypes.STRING,
      field: 'org_id',
      allowNull: false
    },
    spaceId: {
      type: DataTypes.STRING,
      field: 'space_id',
      allowNull: false
    },
  }, {
    timestamps: false,
    freezeTableName: true
  });

};

module.exports = ServiceInstance;