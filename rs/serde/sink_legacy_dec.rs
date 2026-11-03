mod sink_legacy_dec {
    const PRECISION: usize = 18;
    // LegacyDec bounds the atomic magnitude below 2^256 * 10^18.
    const MAX_ATOMS: &str = concat!(
        "115792089237316195423570985008687907853269984665640564039457584007913129639935",
        "999999999999999999"
    );

    fn sign(value: &str) -> (bool, &str) {
        if let Some(digits) = value.strip_prefix('-') {
            (true, digits)
        } else {
            (false, value.strip_prefix('+').unwrap_or(value))
        }
    }

    fn digits(value: &str) -> bool {
        !value.is_empty() && value.bytes().all(|byte| byte.is_ascii_digit())
    }

    fn magnitude(value: &str) -> Result<&str, &'static str> {
        if !digits(value) {
            return Err("reporters_threshold requires decimal digits");
        }
        let value = value.trim_start_matches('0');
        if value.len() > MAX_ATOMS.len() || (value.len() == MAX_ATOMS.len() && value > MAX_ATOMS) {
            return Err("reporters_threshold exceeds the LegacyDec range");
        }
        Ok(if value.is_empty() { "0" } else { value })
    }

    pub(super) fn from_json(value: &str) -> Result<String, &'static str> {
        let (negative, value) = sign(value);
        let (whole, fraction) = match value.split_once('.') {
            Some((whole, fraction)) if digits(fraction) => (whole, fraction),
            Some(_) => return Err("reporters_threshold requires a decimal fraction"),
            None => (value, ""),
        };
        if !digits(whole) {
            return Err("reporters_threshold requires decimal digits");
        }
        if fraction.len() > PRECISION {
            return Err("reporters_threshold exceeds 18 decimal places");
        }

        let atoms = format!("{whole}{fraction:0<PRECISION$}");
        let atoms = magnitude(&atoms)?;
        let sign = if negative && atoms != "0" { "-" } else { "" };
        Ok(format!("{sign}{atoms}"))
    }

    pub(super) fn to_json(value: &str) -> Result<String, &'static str> {
        let (negative, value) = sign(value);
        if value.len() > 1 && value.starts_with('0') {
            return Err("reporters_threshold requires an unambiguous atomic integer");
        }
        let atoms = magnitude(value)?;
        let sign = if negative && atoms != "0" { "-" } else { "" };
        if atoms.len() <= PRECISION {
            return Ok(format!("{sign}0.{atoms:0>PRECISION$}"));
        }

        let (whole, fraction) = atoms.split_at(atoms.len() - PRECISION);
        Ok(format!("{sign}{whole}.{fraction}"))
    }
}
