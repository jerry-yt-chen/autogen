#!/bin/sh

# Get App flags decrypted by macgyver
getAppFlags() {
  # Compose decrypt command flags
  none="none"
  macgyverFlags="--flags=\"$FLAGS\" --keysType=text"
  output=$FLAGS
  # crypto provider
  if [ "${PLATFORM}" == "gcp" ];then
    if [ $kms_cryptoProvider ] && [ $kms_cryptoProvider != $none ]; then
      macgyverFlags="$macgyverFlags --cryptoProvider=$kms_cryptoProvider"
    fi
    # GCPprojectID
    if [ $kms_GCPprojectID ] && [ $kms_GCPprojectID != $none ]; then
      macgyverFlags="$macgyverFlags --GCPprojectID=$kms_GCPprojectID"
    fi
    # GCPlocationID
    if [ $kms_GCPlocationID ] && [ $kms_GCPlocationID != $none ]; then
      macgyverFlags="$macgyverFlags --GCPlocationID=$kms_GCPlocationID"
    fi
    # GCPkeyRingID
    if [ $kms_GCPkeyRingID ] && [ $kms_GCPkeyRingID != $none ]; then
      macgyverFlags="$macgyverFlags --GCPkeyRingID=$kms_GCPkeyRingID"
    fi
    # GCPcryptoKeyID
    if [ $kms_GCPcryptoKeyID ] && [ $kms_GCPcryptoKeyID != $none ]; then
      macgyverFlags="$macgyverFlags --GCPcryptoKeyID=$kms_GCPcryptoKeyID"
    fi
  elif [ "${PLATFORM}" == "aws" ];then
   # AWSlocationID
   if [ $kms_AWSlocationID ] && [ $kms_AWSlocationID != $none ]; then
     macgyverFlags="$macgyverFlags --AWSlocationID=$kms_AWSlocationID"
   fi
   # AWScryptoKeyID
   if [ $kms_AWScryptoKeyID ] && [ $kms_AWScryptoKeyID != $none ]; then
     macgyverFlags="$macgyverFlags --AWScryptoKeyID=$kms_AWScryptoKeyID"
   fi
  fi

  if [[ "${PLATFORM}" == "gcp" || "${PLATFORM}" == "aws"  ]];then
    # Execute decrypt command
    decryptCmd="./macgyver decrypt $macgyverFlags"
    output=$(eval $decryptCmd)
    if [[ $? -ne 0 ]]; then
      echo >&2 "Decrypting flags is failed."
      exit 1
    fi
  fi
  echo $output
}

# Catch SIGTERM
_term() {
  # Send SIGTERM to child
  # And keep check child process is running
  kill -TERM "$child"
  stoped=0
  sec=0
  while [ $stoped -eq 0 ]
  do
    sec=$((sec + 1))
    if [ ! -e /proc/$child ]; then
      stoped=1
    else
      if [ $sec -eq 60 ]; then
        echo term timeout 60 sec
      fi
      sleep 1
    fi
  done
  # wait is used to capture the return code
  wait $child
  exit_status=$?
  echo graceful shut down with code $exit_status
  exit 0
}
trap _term TERM

# Get flags
appFlags=$(getAppFlags)
if [[ $? -ne 0 ]]; then
  exit 1
fi

# Run app with flags environment variable
cmd="${WORK_DIR}/main $appFlags"
echo -----------------------------------------------------
echo $(date)
$cmd &
child=$!
# first wait will be interrupted be a signal
wait $child
# second wait is used to capture the return code
wait $child
exit_status=$?
echo end of child process with code $exit_status
# send abnormal exit status to /logs for monitoring
exit $exit_status
