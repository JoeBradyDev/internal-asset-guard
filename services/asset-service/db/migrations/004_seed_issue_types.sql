-- UP

--- 1. CONFIGURATION DRIFT (System State & Forensic Changes) ---
INSERT INTO issue_type (category_id, name, description) VALUES
(1, 'Firewall Rule Change', 'Unauthorized modification to local or network firewall rules.'),
(1, 'Registry Key Modification', 'Critical security-related registry keys differ from baseline.'),
(1, 'Unsigned Driver Installation', 'A new driver was installed that lacks a valid digital signature.'),
(1, 'Local Admin Group Expansion', 'New users added to the local Administrators group.'),
(1, 'Disabled Security Agent', 'The EDR, AV, or DLP agent service has been stopped or disabled.'),
(1, 'Windows Update Service Disabled', 'The automatic update service has been turned off.'),
(1, 'Unauthorized Browser Extension', 'A high-risk or blacklisted browser extension was detected.'),
(1, 'Host File Modification', 'The system hosts file has been altered to redirect traffic.'),
(1, 'SSH Authorized_Keys Change', 'New entries detected in a user''s authorized_keys file.'),
(1, 'BitLocker Decryption Started', 'A drive previously encrypted is currently being decrypted.'),
(1, 'Audit Policy Disabled', 'Local security auditing policies were turned off or reduced.'),
(1, 'PowerShell Execution Policy Lowered', 'Policy changed from Restricted to Bypass or Unrestricted.'),
(1, 'New Scheduled Task Created', 'A non-standard task was registered to run with SYSTEM privileges.'),
(1, 'Unexpected Kernel Module', 'A non-standard LKM (Linux Kernel Module) was loaded.'),
(1, 'Print Spooler Enabled', 'The Print Spooler service is active on a high-value server (RCE risk).'),
(1, 'Guest Account Password Set', 'The built-in guest account was enabled and assigned a password.'),
(1, 'DNS Resolver Overridden', 'Local DNS settings have been changed to point to unauthorized resolvers.'),
(1, 'New Root CA Installed', 'A new root certificate was added to the system trusted store.');

--- 2. VULNERABILITY CLASSES (Generic buckets for CVE grouping) ---
INSERT INTO issue_type (category_id, name, description) VALUES
(2, 'Remote Code Execution (RCE)', 'Flaw allowing an attacker to execute arbitrary commands on the host.'),
(2, 'SQL Injection (SQLi)', 'Input validation flaw allowing unauthorized database queries.'),
(2, 'Cross-Site Scripting (XSS)', 'Injection of malicious scripts into trusted websites.'),
(2, 'Buffer Overflow', 'Memory handling error that can lead to execution hijacking.'),
(2, 'Path Traversal', 'Flaw allowing access to files outside of the intended directory.'),
(2, 'Broken Authentication', 'Flaws in session management or credential validation.'),
(2, 'Insecure Deserialization', 'Untrusted data used to abuse logic or execute code.'),
(2, 'Server-Side Request Forgery (SSRF)', 'Abusing the server to make requests to internal resources.'),
(2, 'Cryptographic Failure', 'Use of weak hashes, expired certs, or cleartext sensitive data.'),
(2, 'Privilege Escalation', 'Flaw allowing a low-level user to gain root or admin access.'),
(2, 'Outdated Kernel', 'The OS kernel is no longer receiving security patches.'),
(2, 'End-of-Life (EOL) Software', 'Software version is no longer supported by the vendor.'),
(2, 'Information Disclosure', 'Flaw resulting in the exposure of sensitive system data.'),
(2, 'Insecure API Endpoint', 'API lacks proper rate limiting, authentication, or input filtering.'),
(2, 'XML External Entity (XXE)', 'Processing of XML input containing a reference to an external entity.'),
(2, 'Unpatched Zero-Day', 'Vulnerability being actively exploited with no official patch available.');

--- 3. NON-COMPLIANCE (Hardening, Benchmarks & Cloud Governance) ---
INSERT INTO issue_type (category_id, name, description) VALUES
(3, 'Password Complexity Violation', 'Fails to enforce minimum character or complexity requirements.'),
(3, 'Disk Encryption Disabled', 'Full disk encryption (BitLocker/FileVault) is not active.'),
(3, 'SSH Root Login Enabled', 'Remote root access is permitted via SSH.'),
(3, 'Legacy TLS Active', 'Support for TLS 1.0 or 1.1 is still enabled.'),
(3, 'Unattended Sleep Timeout', 'Screen lock timeout exceeds the policy limit.'),
(3, 'SMBv1 Enabled', 'Insecure legacy file sharing protocol is active.'),
(3, 'HTTP Management Access', 'Device management interface is accessible over unencrypted HTTP.'),
(3, 'Insecure S3 Bucket Policy', 'Cloud storage bucket allows public read or write access.'),
(3, 'Unrestricted Inbound SSH', 'Cloud security group allows SSH from any IP (0.0.0.0/0).'),
(3, 'IMDSv1 Enabled', 'Cloud instance using legacy metadata service (SSRF risk).'),
(3, 'Docker Socket Mounted', 'Container has direct access to host Docker daemon.'),
(3, 'Privileged Container', 'Pod/Container running with elevated kernel capabilities.'),
(3, 'Unencrypted RDS Instance', 'Cloud database instance is not encrypted at rest.'),
(3, 'Publicly Accessible Database', 'Database port (3306, 5432, etc.) is open to the internet.'),
(3, 'Kubernetes Dashboard Exposed', 'The K8s management UI is accessible without authentication.'),
(3, 'Missing Security Contact', 'Cloud account lacks a registered security emergency contact.'),
(3, 'Flow Logs Disabled', 'Network traffic logging is not enabled for the VPC/Subnet.'),
(3, 'Default VPC in Use', 'Resources are deployed in the unhardened default cloud network.');

--- 4. UNAUTHORIZED ASSET (Detection & Network Hygiene) ---
INSERT INTO issue_type (category_id, name, description) VALUES
(4, 'Rogue Wireless Access Point', 'Non-sanctioned Wi-Fi hotspot detected on the network.'),
(4, 'Unknown MAC Address', 'Device with an unrecognized OUI connected to a physical port.'),
(4, 'Unauthorized Virtual Machine', 'VM detected that is not in the central inventory.'),
(4, 'Shadow IT SaaS App', 'Unauthorized cloud storage or file-sharing application in use.'),
(4, 'Rogue DHCP Server', 'Unauthorized device attempting to lease IP addresses.'),
(4, 'Cryptocurrency Miner', 'System detected communicating with known mining pools.'),
(4, 'Tor Exit Node Entry', 'Asset is routing traffic through the Tor network.'),
(4, 'Unauthorized USB Storage', 'Mass storage device connected to a restricted endpoint.'),
(4, 'Double-Homed Host', 'Asset connected to both corporate and untrusted networks.'),
(4, 'Rogue VPN Gateway', 'Unauthorized VPN server detected within the perimeter.'),
(4, 'Unauthorized IoT Sensor', 'Smart camera or sensor found on a production VLAN.'),
(4, 'Unmanaged Network Switch', 'Dumb switch connected to a managed access port.'),
(4, 'Beaconing Behavior', 'Rhythmic outbound traffic to an unclassified external IP.');

--- 5. ACCOUNT ISSUE (Identity, IAM & Privilege) ---
INSERT INTO issue_type (category_id, name, description) VALUES
(5, 'Dormant Account', 'User account has not logged in for over 90 days.'),
(5, 'MFA Disabled', 'Privileged user account lacks Multi-Factor Authentication.'),
(5, 'Account Password Never Expires', 'User account set to override the rotation policy.'),
(5, 'Excessive Admin Privileges', 'User assigned Domain Admin rights without a change request.'),
(5, 'Service Account Interactive Login', 'Non-human account used for interactive console login.'),
(5, 'Orphaned Account', 'Account remains active for a user no longer in the HR system.'),
(5, 'Password Spraying Target', 'Account has experienced 50+ failed login attempts.'),
(5, 'Shared Credentials Detected', 'Account logged in from two distant locations simultaneously.'),
(5, 'Static AWS Access Keys', 'Long-lived IAM access keys have not been rotated in 180 days.'),
(5, 'Default Admin Credentials', 'Device or application using factory default login/password.'),
(5, 'Excessive IAM Permissions', 'IAM role has wildcard (*) permissions to sensitive services.'),
(5, 'Access Key Leaked in Code', 'Credentials found in a public or internal code repository.'),
(5, 'Concurrent Privileged Sessions', 'Multiple admin sessions from different source IPs on one account.'),
(5, 'Password in Description Field', 'Sensitive information found in AD or IAM metadata fields.'),
(5, 'Lack of Least Privilege', 'User account has write access to a read-only data store.'),
(5, 'Unused IAM Role', 'Cloud role has not been assumed or used in over 60 days.');
