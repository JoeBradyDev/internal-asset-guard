-- UP

-- CIS v8.1 Asset Classes
INSERT INTO cis_asset_class (name, definition) VALUES
('Devices', 'Hardware assets: end-user devices, servers, IoT, and mobile.'),
('Software', 'Operating systems, applications, and specialized software.'),
('Data', 'Information assets: sensitive, log, and operational data.'),
('Users', 'Workforce accounts, service accounts, and administrators.'),
('Network', 'Network infrastructure: routers, switches, and firewalls.'),
('Documentation', 'Policies, procedures, and security diagrams.');

INSERT INTO cis_asset_criticality (name, value) VALUES
('Low', 1), ('Medium', 2), ('High', 3), ('Critical', 4);

INSERT INTO asset_source (name) VALUES
('CrowdStrike'), ('Qualys'), ('SolarWinds'), ('vCenter'), ('Active Directory'), ('Manual');

INSERT INTO issue_source (name) VALUES
('Tenable'), ('Rapid7'), ('Cisco Advisory'), ('GitHub Dependabot'), ('Manual Audit'), ('NIST NVD');

INSERT INTO cis_issue_status (name) VALUES
('New'), ('Confirmed'), ('In Progress'), ('Remediated'), ('Resolved'), ('False Positive'), ('Risk Accepted');

-- Operational Issue Categories
INSERT INTO cis_issue_category (name, description) VALUES
('Configuration Drift', 'A setting that was previously secure but has changed over time.'),
('Vulnerability', 'A known flaw in software or hardware (often tracked via CVEs).'),
('Non-Compliance', 'An asset that fails to meet a specific CIS Benchmark recommendation.'),
('Unauthorized Asset', 'An asset found on the network that is not in the authorized inventory.'),
('Account Issue', 'Problems related to user identity, privileges, or dormant accounts.');
